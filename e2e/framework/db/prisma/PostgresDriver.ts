import { PrismaClient } from "@prisma/client";
import { DBDriver } from "../DBDriver";
import { SpecDataManager } from "../gauge-ts/data-manager/SpecDataManager";
import { SpecDBEnvInfo } from "../definitions";

const prisma = new PrismaClient();

class PostgresDriver implements DBDriver {
    async read(tableName: string, where: object): Promise<any> {
        const client = this.createClient(tableName);
        const res = await client.findMany({ where: where })
        return res
    }

    async create(tableName: string, data: object): Promise<any> {
        const client = this.createClient(tableName);
        const res = await client.create({data: data});
        SpecDataManager.put(SpecDBEnvInfo, res);
        return res
    }

    private createClient(tableName: string) {
        return (prisma as any)[tableName];
    }
}
