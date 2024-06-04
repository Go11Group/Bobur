-------- 1- task   --------
create index index_user on user1 (name, fname, phone);

explain (analyse)
select * from user1 where name = 'Alta';

-- 1 marta ishlatganimda 

Bitmap Heap Scan on user1  (cost=5.10..293.53 rows=88 width=37) (actual time=0.147..0.198 rows=77 loops=1)
   Recheck Cond: ((name)::text = 'Alta'::text)
   Heap Blocks: exact=76
   ->  Bitmap Index Scan on index_user  (cost=0.00..5.08 rows=88 width=0) (actual time=0.018..0.018 rows=77 loops=1)
         Index Cond: ((name)::text = 'Alta'::text)
 Planning Time: 0.074 ms
 Execution Time: 0.300 ms
 
 ----------------------------------------

-- 2 marta ishlatganimda

Bitmap Heap Scan on user1  (cost=5.10..293.53 rows=88 width=37) (actual time=0.061..0.179 rows=77 loops=1)
   Recheck Cond: ((name)::text = 'Alta'::text)
   Heap Blocks: exact=76
   ->  Bitmap Index Scan on index_user  (cost=0.00..5.08 rows=88 width=0) (actual time=0.036..0.036 rows=77 loops=1)
         Index Cond: ((name)::text = 'Alta'::text)
 Planning Time: 0.112 ms
 Execution Time: 0.212 ms

-------------------------------------

-- 3 marta ishlatganimda

                                    QUERY PLAN                                                      
---------------------------------------------------------------------------------------------------------------------
 Bitmap Heap Scan on user1  (cost=5.10..293.53 rows=88 width=37) (actual time=0.055..0.167 rows=77 loops=1)
   Recheck Cond: ((name)::text = 'Alta'::text)
   Heap Blocks: exact=76
   ->  Bitmap Index Scan on index_user  (cost=0.00..5.08 rows=88 width=0) (actual time=0.032..0.033 rows=77 loops=1)
         Index Cond: ((name)::text = 'Alta'::text)
 Planning Time: 0.094 ms
 Execution Time: 0.196 ms


drop index index_user

----------- 2 - task -----------
create index index_user on user1 (fname, name, phone);
explain (analyse)
select * from user1 where name = 'Alta';

-- 1 marta ishlatganimda

                         QUERY PLAN                                                     
--------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.339..7.576 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.104..4.655 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.110 ms
 Execution Time: 7.611 ms

-- 2 marta ishlatganimda

                        QUERY PLAN                                                      
---------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.513..16.467 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.293..10.814 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.133 ms
 Execution Time: 16.494 ms

-- 3 marta ishlatganimda

                         QUERY PLAN                                                     
--------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.479..10.673 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.127..6.013 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.106 ms
 Execution Time: 10.697 ms



drop index index_user




------------ 3 - task -----------

create index index_user on user1 (fname, phone, name);
explain (analyse)
select * from user1 where name = 'Alta';


-- 1 marta ishlatganimda

                        QUERY PLAN                                                     
--------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.208..7.709 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.097..4.613 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.080 ms
 Execution Time: 7.721 ms

-- 2 marta ishlatganimda

                        QUERY PLAN                                                     
--------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.501..11.164 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.170..6.496 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.105 ms
 Execution Time: 11.189 ms


-- 3 marta ishlatganumda

                            QUERY PLAN                                                     
--------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..4750.74 rows=88 width=37) (actual time=0.461..14.318 rows=77 loops=1)
   Workers Planned: 1
   Workers Launched: 1
   ->  Parallel Seq Scan on user1  (cost=0.00..3741.94 rows=52 width=37) (actual time=0.419..9.171 rows=38 loops=2)
         Filter: ((name)::text = 'Alta'::text)
         Rows Removed by Filter: 133238
 Planning Time: 0.137 ms
 Execution Time: 14.343 ms


drop index index_user


---------------------------------------------------------------------------------------------------------------------------

                    -- 2 - task -- 
create index index_user on user1 (name, fname);

explain (analyse)
select * from user1 where name = 'Jacky' and fname = 'Kuhn';


                        QUERY PLAN                                                     
-------------------------------------------------------------------------------------------------------------------
 Index Scan using index_user on user1  (cost=0.42..8.44 rows=1 width=37) (actual time=0.032..0.036 rows=2 loops=1)
   Index Cond: (((name)::text = 'Jacky'::text) AND ((fname)::text = 'Kuhn'::text))
 Planning Time: 0.097 ms
 Execution Time: 0.058 ms

drop index index_user;

------------------ 2 - masala--------------------------------------
create index index_user on user1 (name, fname);

explain (analyse)
select * from user1 where name = 'Jacky' and fname = 'Kuhn';

                        QUERY PLAN                                                     
-------------------------------------------------------------------------------------------------------------------
 Index Scan using index_user on user1  (cost=0.42..8.44 rows=1 width=37) (actual time=0.024..0.027 rows=2 loops=1)
   Index Cond: (((name)::text = 'Jacky'::text) AND ((fname)::text = 'Kuhn'::text))
 Planning Time: 0.093 ms
 Execution Time: 0.046 ms
