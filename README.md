# Jubo 測試後端

## 環境建置
1. docker-compose up -d, 這一步會啟動db 跟後端本體
2. 這部開始處理db schema 跟 seed data
2. docker cp db/migrations/000001_init_schema.up.sql db:/000001_init_schema.up.sql
3. docker cp db/migrations/20240123063149_seed_file.up.sql db:/20240123063149_seed_file.up.sql
4. docker exec -it db bash
5. psql -U postgres -d postgres -f 000001_init_schema.up.sql
6. psql -U postgres -d postgres -f 20240123063149_seed_file.up.sql

## api 內容

#### 詳見test.http
