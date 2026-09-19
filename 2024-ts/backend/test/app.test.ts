import request from "supertest";
import { initApp } from "../src/app";
import {closeDatabase} from "../src/db";

describe("GET /api/hello", () => {
  afterAll(async () => {
    await closeDatabase();
  });

  it("returns a hello message", async () => {
    const app = await initApp();
    const response = await request(app)
      .get("/api/hello");

    expect(response.status).toBe(200);

    expect(response.body).toEqual({
      message: "Hello!",
      tasks: [],
    });
  });
});