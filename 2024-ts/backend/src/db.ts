import { MongoClient, Db } from "mongodb";

const uri = "mongodb://admin:password@localhost:27017/tasks?authSource=admin";
const client = new MongoClient(uri, {
  serverSelectionTimeoutMS: 2000,
});

let db: Db;

export async function connectToDatabase(): Promise<Db> {
  await client.connect();
  db = client.db("tasks");

  return db;
}

export function getDatabase(): Db {
  if (!db) {
    throw new Error("Database not connected");
  }

  return db;
}

export async function closeDatabase(): Promise<void> {
  await client.close();
}