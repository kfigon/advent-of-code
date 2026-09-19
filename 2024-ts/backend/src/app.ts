import express from "express";
import { connectToDatabase, getDatabase } from "./db";

export async function initApp() {
  const app = express();
  app.use(express.json());
  
  await connectToDatabase();

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