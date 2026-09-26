import { MongoClient, Db } from "mongodb";
import { DatabaseConfig } from "./config";


let client: MongoClient
let db: Db;

export async function connectToDatabase({
    user,
    password,
    port,
    tableName,
  }: DatabaseConfig): Promise<Db> {

  const uri = `mongodb://${user}:${password}@localhost:${port}/${tableName}?authSource=admin`;
  client = new MongoClient(uri, {
    serverSelectionTimeoutMS: 2000,
  });
  await client.connect();
  db = client.db(tableName);

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