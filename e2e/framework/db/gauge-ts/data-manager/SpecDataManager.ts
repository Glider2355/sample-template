import { DataStoreFactory } from "gauge-ts"

export const SpecDataManager = {
    get: (key: string): any => {
        const dataStore = DataStoreFactory.getSpecDataStore();
        return dataStore.get(key);
    },
    put: (key: string, value: any): void => {
        const dataStore = DataStoreFactory.getSpecDataStore();
        dataStore.put(key, value);
    }
}
