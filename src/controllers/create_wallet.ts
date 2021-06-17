import express from "express";
import crypto from "crypto-js";
import { calcSignature } from "../util/signature";
import axios from "axios";
import { db } from "../util/firebase";

require("dotenv").config();

export const createWallet = async (
  _req: express.Request,
  _res: express.Response,
) => {
  // get data from req
  const data = _req.body;
  // get user id from req header
  const userID = await _req.header("user_id");

  const accessKey = process.env.RAPYD_ACCESS_KEY!;
  const secretKey = process.env.RAPYD_SECRET_KEY!;
  const salt = crypto.lib.WordArray.random(12).toString();
  const timestamp = (Math.floor(new Date().getTime() / 1000) - 10).toString();
  const signature = calcSignature(
    "post",
    "/v1/user",
    salt,
    accessKey,
    secretKey,
    JSON.stringify(data),
  );
  try {
    const result: any = await axios.post(
      "https://sandboxapi.rapyd.net/v1/user",
      data,
      {
        headers: {
          "content-type": "application/json",
          access_key: accessKey,
          salt: salt,
          timestamp: timestamp,
          signature: signature,
        },
      },
    );
    await _res.json({
      data: result.data,
      message: "Successfully created",
    });
    if (!userID) {
      _res.send("there was an error storing data in db");
      return;
    }
    console.log(result.data.data.id);
    try {
      await db.collection("users").doc(userID).update({
        ewalletID: result.data.data.id,
      });
    } catch (err) {
      console.log(err);
    }
  } catch (err) {
    await _res.status(400).json({
      error: err,
      message: "Invalid data passed",
    });
  }
};
