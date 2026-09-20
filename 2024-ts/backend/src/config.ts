export type Config = {
    dbConf: DatabaseConfig
    appConf: AppConfig
};

export type AppConfig = {
    port: number
}

export type DatabaseConfig = {
    user: string
    password: string
    port: number
    tableName: string
}

function requiredEnv(name: string): string {
    const v = process.env[name]    
    if (!v) throw new Error(`missing env var ${name}`)

    return v;
}

function readEnvWithDefault(name: string, defaultV: string): string {
    return process.env[name] ?? defaultV;
}

function requiredEnvAs<T>(name: string, fn: (v: string)=>T): T {
    return fn(requiredEnv(name));
}

export function readConfigFromEnv(): Config{
    const toBool = (v: string) => {
        switch(v) {
        case "true": return true
        case "false": return false            
        default: throw new Error(`invalid boolean value ${v}`)
        }
    }

    return {
        dbConf: {
           password: requiredEnv("DATABASE_PASS"),
           user: requiredEnv("DATABASE_USER"),
           tableName: requiredEnv("DATABASE_TABLE"),
           port: requiredEnvAs("DATABASE_PORT", Number),
        },
        appConf: {
           port: Number(readEnvWithDefault("SERVER_PORT", "3000"))
        }
    }
}
