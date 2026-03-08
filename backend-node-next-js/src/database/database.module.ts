import { Global, Module } from '@nestjs/common';
import Database from 'better-sqlite3';

export const DB_TOKEN = 'DATABASE';

@Global()
@Module({
  providers: [
    {
      provide: DB_TOKEN,
      useFactory: () => {
        const db = new Database('aerostack.db');
        db.pragma('journal_mode = WAL');
        db.pragma('foreign_keys = ON');
        return db;
      },
    },
  ],
  exports: [DB_TOKEN],
})
export class DatabaseModulo {}
