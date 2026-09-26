import express from "express";
import path from "node:path";
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

  // serve frontend
  const frontendPath = path.resolve(__dirname, "../frontend/dist");
  app.use(express.static(frontendPath));
  app.get("/{*splat}", (_req, res) => {
    res.sendFile(path.join(frontendPath, "index.html"));
  });

  return app;
}