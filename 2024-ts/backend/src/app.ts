import express from "express";
import { connectToDatabase, getDatabase } from "./db";

export async function initApp() {
   console.log("starting") 
  const app = express();

  app.use(express.json());
  
   console.log("connecting to db") 
  await connectToDatabase();

   console.log("connected to db") 
  app.get("/api/hello", async (req, res) => {
    try {
      const db = getDatabase();

      console.log("Got database");

      const tasks = await db
        .collection("tasks")
        .find()
        .toArray();

      console.log("Got tasks:", tasks);

      res.json({
        message: "Hello!",
        tasks,
      });
    } catch (error) {
      console.error("DATABASE ERROR:", error);

      res.status(500).json({
        error: "Database error",
      });
    }
  });

  return app;
}