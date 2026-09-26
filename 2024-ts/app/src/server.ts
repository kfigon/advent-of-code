import { initApp } from "./app";
import { readConfigFromEnv } from "./config";
import dotenv from "dotenv"

dotenv.config();

async function startServer() {

  const conf = readConfigFromEnv();
  const app = await initApp(conf);

  const port = conf.appConf.port;
  app.listen(port, () => {
    console.log(`Server listening on http://localhost:${port}`);
  });
}

startServer();