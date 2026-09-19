import { initApp } from "./app";

const PORT = 3000;

async function startServer() {
  const app = await initApp();

  app.listen(PORT, () => {
    console.log(`Server listening on http://localhost:${PORT}`);
  });
}

startServer();