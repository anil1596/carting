import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'foodies', host = 'localhost', user = 'anil', password = '205474')
cursor = postgres.cursor()


# TABLE VENDORS
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'vendors')
""")

presence = cursor.fetchone()[0]
print(presence)

if presence:
    print('vendors table exists, deleting here')
    cursor.execute("""
        drop table vendors
    """)

#                       CREATE TABLE VENDORS                                         #              
cursor.execute("""
    create table vendors(
        vendorid bigserial PRIMARY KEY,
        owner varchar(60) NOT NULL,
        vendorname varchar(80) NOT NULL,
        email varchar(35) NOT NULL,
        mobile varchar(12)[] NOT NULL,
        address text NOT NULL,
        offer text,
        imageaddress varchar(50),
        password varchar(40) NOT NULL
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


# TABLE ITEMSMENU
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'itemmenu')
""")

presence = cursor.fetchone()[0]
print(presence)

if presence:
    print('ItemsMenu table exists, deleting here')
    cursor.execute("""
        drop table itemmenu
    """)
#                       CREATE TABLE ITEMSMENU                                      #              
cursor.execute("""
    create table itemmenu(
        item_no bigserial PRIMARY KEY,
        vendor_id int NOT NULL,
        item_name varchar(60) NOT NULL,
        item_type character(10) NOT NULL CHECK(item_type in ('starter','main','desert')) ,
        item_nature character(1) NOT NULL CHECK(item_nature in ('v','n')),
        price money NOT NULL,
        item_description text NOT NULL,
        offer text,
        imageaddress varchar(50)
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                               CHECK FOR TABLE CUSTOMERS                                #
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'cutomers')
""")

presence = cursor.fetchone()[0]
print(presence)


 #                       DELETE CUSTOMERS          IF EXISTS                            #
if presence:
    print('cutomers table exists, deleting here')
    cursor.execute("""
        drop table cutomers
    """)

#                       CREATE TABLE CUSTOMERS                                            #
cursor.execute("""
    create table cutomers(
            cutomer_id bigserial PRIMARY KEY NOT NULL,
            cutomer_name varchar(80) NOT NULL,
            emailid varchar(50)[] NOT NULL,
            mobile varchar(12)[] NOT NULL,
            address text,
            password varchar(40) NOT NULL
    )
""")
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                               CHECK FOR TABLE FOODIES_RECORD                      #
# cursor.execute("""
#    select exists(select 0 from pg_class where relname = 'foodiesrecord')
# """)

# presence = cursor.fetchone()[0]
# print(presence)

#  #                       DELETE FOODIESRECORD IF EXISTS                              #
# if presence:
#     print('foodiesrecord table exists, deleting here')
#     cursor.execute("""
#         drop table foodiesrecord
#     """)

# #                       CREATE TABLE FOODIESRECORD                                     #
# cursor.execute("""
#     create table foodiesrecord(
#             transactionid bigserial PRIMARY KEY NOT NULL,
#             ordered_on timestamp NOT NULL,
#             customer_id int NOT NULL,
#             delivered_on timestamp,
#             vendorid int NOT NULL,
#             status varchar(20) NOT NULL,
#             bill money NOT NULL
#     )
# """)
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)


#                               CHECK FOR TABLE ACCOUNTS_RECORD                       #
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'accounts_record')
""")
# it returns true if there is at least row in database and it will be a tuple
presence = cursor.fetchone()[0]
print(presence)


 #                       DELETE FOODIES_RECORD IF EXISTS                              #
if presence:
    print('accounts_record table exists, deleting here')
    cursor.execute("""
        drop table accounts_record
    """)

#                       CREATE TABLE ACCOUNTS_RECORD                                    #
cursor.execute("""
    create table accounts_record(
        v_id bigserial PRIMARY KEY NOT NULL,
        account smallint NOT NULL
    )
""")
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


# TABLE ORDERsRECORD
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'ordersrecord')
""")

presence = cursor.fetchone()[0]
print(presence)

if presence:
    print(' table exists, deleting here')
    cursor.execute("""
        drop table ordersrecord
    """)


#                       CREATE TABLE ORDERsRECORD                                   #              
cursor.execute("""
    create table ordersrecord(
        order_id int PRIMARY KEY,
        vendor_id int NOT NULL,
        cutomer_id int NOT NULL,        
        ordered_placed_on timestamp NOT NULL,
        ordered_deliverd_on timestamp ,
        order_status character(1) CHECK(order_status in('y','n')),
        description text NOT NULL,
        amount money NOT NULL
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


