import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'web_portal', host = 'localhost', user = 'postgres', password = 'anil205474')
cursor = postgres.cursor()

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'admin')
""")

presence = cursor.fetchone()[0]

if presence:
    print('admin table exists')
    cursor.execute("""
        drop table admin
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'students')
""")

presence = cursor.fetchone()[0]

if presence:
    print('students table exists')
    cursor.execute("""
        drop table students
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
        roll_no text,
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
        password text,
        PRIMARY KEY(item_no)
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


# #                       CREATE TABLE CUSTOMER                                         #
# cursor.execute("""
#     create table customer(
#         customer_id bigserial PRIMARY KEY,
#         name text NOT NULL,
#         email text NOT NULL,
#         mobile text,
#         hostel text NOT NULL,
#         room text NOT NULL,
#         bid_amount text NOT NULL,
#         item_id int NOT NULL references item(item_no)
#     )
# """)

# #                       COMMIT AND ROLLBACK IF EXCEPTION                             #
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)


#                       CREATE TABLE STUDENTS                                      #
cursor.execute("""
    create table students(
        roll_no text NOT NULL ,
        password text NOT NULL,
        block boolean NOT NULL,
        total int default 0,
        PRIMARY KEY(roll_no)
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)

#                       CREATE TABLE ADMINS                                      #
cursor.execute("""
    create table admin(
        employee_no text NOT NULL ,
        password text NOT NULL,
        PRIMARY KEY(employee_no)
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)

