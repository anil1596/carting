import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'web_portal', host = 'localhost', user = 'anil', password = 'anil205474')
cursor = postgres.cursor()

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'customer')
""")

presence = cursor.fetchone()[0]

if presence:
    print('customer table exists')
    cursor.execute("""
        drop table customer
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'item')
""")

presence = cursor.fetchone()[0]
if presence:
    print('item table exists')
    cursor.execute("""
        drop table item
    """)





try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)



#                       CREATE TABLE ITEM                                        #
cursor.execute("""
    create table item(
        item_no bigserial  ,
        name text ,
        email text ,
        mobile text ,
        hostel text ,
        room text ,
        item_name text ,
        item_type text ,
        sold boolean ,
        price text,
        item_description text ,
        imageaddress text,
        PRIMARY KEY(item_no)
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                       CREATE TABLE CUSTOMER                                         #
cursor.execute("""
    create table customer(
        customer_id bigserial PRIMARY KEY,
        name text NOT NULL,
        email text NOT NULL,
        mobile text,
        hostel text NOT NULL,
        room text NOT NULL,
        bid_amount text NOT NULL,
        item_id int NOT NULL references item(item_no)
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


# #                       CREATE TABLE ITEMS                                      #
# cursor.execute("""
#     create table item(
#         seller_id int NOT NULL references seller(seller_id),
#         item_no bigserial  ,
#         item_name text NOT NULL,
#         item_type varchar(10) NOT NULL CHECK(item_type in ('starter','main','desert')) ,
#         sold boolean NOT NULL ,
#         price text NOT NULL,
#         item_description text NOT NULL,
#         imageaddress text,
#         discount double precision,
#         PRIMARY KEY(seller_id, item_no)
#     )
# """)
# #                       COMMIT AND ROLLBACK IF EXCEPTION                             #
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)


