import express from "express";
import { addFunds } from "../controllers/add_funds";

const router = express.Router();

router.post("/afunds", addFunds);

export default module.exports = { router };
