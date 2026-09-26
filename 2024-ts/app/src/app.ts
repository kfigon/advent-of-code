import express from "express";
import { connectToDatabase, getDatabase } from "./db";
import {Config} from "./config"

export async function initApp(conf: Config) {
  const app = express();
  app.use(express.json());
  
  await connectToDatabase(conf.dbConf);

  app.get("/api/hello", async (req, res) => {
    const db = getDatabase();

    const tasks = await db
      .collection("tasks")
      .find()
      .toArray();

    res.json({
      message: "Hello!",
      tasks,
    });
  });

  return app;
}