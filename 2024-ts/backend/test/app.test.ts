import request from "supertest";
import { initApp } from "../src/app";
import {closeDatabase} from "../src/db";

let app: Awaited<ReturnType<typeof initApp>>;

beforeEach(async () => {
  app = await initApp();
});

afterAll(async () => {
  await closeDatabase();
});

describe("GET /api/hello", () => {
  it("returns a hello message", async () => {
    const response = await request(app)
      .get("/api/hello");

    expect(response.status).toBe(200);

    expect(response.body).toEqual({
      message: "Hello!",
      tasks: [],
    });
  });
});