export interface DBDriver {
    create(tableName: string, data: object): Promise<any>;
    read(tableName: string, where: object): Promise<any>;
}

