import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'web_portal', host = 'localhost', user = 'anil', password = 'anil205474')
cursor = postgres.cursor()

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'item')
""")

presence = cursor.fetchone()[0]
if presence:
    print('item table exists')
    cursor.execute("""
        drop table item
    """)

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
   select exists(select 0 from pg_class where relname = 'seller')
""")
    
presence = cursor.fetchone()[0]

if presence:
    print('seller table exists')
    cursor.execute("""
        drop table seller
    """)


try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)



#                       CREATE TABLE SELLER                                         #
cursor.execute("""
    create table seller(
        seller_id bigserial PRIMARY KEY,
        first_name text NOT NULL,
        last_name text NOT NULL,
        email text NOT NULL,
        mobile varchar(12)[] NOT NULL,
        hostel text NOT NULL,
        room text NOT NULL,
        password text NOT NULL
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
        first_name text NOT NULL,
        last_name text NOT NULL,
        email text NOT NULL,
        mobile varchar(12)[] NOT NULL,
        hostel text NOT NULL,
        room text NOT NULL,
        password text NOT NULL
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                       CREATE TABLE ITEMS                                      #
cursor.execute("""
    create table item(
        seller_id int NOT NULL references seller(seller_id),
        item_no bigserial  ,
        item_name text NOT NULL,
        item_type varchar(10) NOT NULL CHECK(item_type in ('starter','main','desert')) ,
        sold boolean NOT NULL ,
        price text NOT NULL,
        item_description text NOT NULL,
        imageaddress text,
        discount double precision,
        PRIMARY KEY(seller_id, item_no)
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)



# #                       CREATE TABLE CUSTOMERS                                            #
# cursor.execute("""
#     create table customers(
#             customer_id bigserial PRIMARY KEY NOT NULL,
#             customer_name text NOT NULL,
#             emailid text NOT NULL,
#             mobile varchar(12)[] NOT NULL,
#             address text,
#             password text NOT NULL
#     )
# """)
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)


# #                               CHECK FOR TABLE FOODIES_RECORD                      #
# # cursor.execute("""
# #    select exists(select 0 from pg_class where relname = 'foodiesrecord')
# # """)

# # presence = cursor.fetchone()[0]
# # print(presence)

# #  #                       DELETE FOODIESRECORD IF EXISTS                              #
# # if presence:
# #     print('foodiesrecord table exists, deleting here')
# #     cursor.execute("""
# #         drop table foodiesrecord
# #     """)

# # #                       CREATE TABLE FOODIESRECORD                                     #
# # cursor.execute("""
# #     create table foodiesrecord(
# #             transactionid bigserial PRIMARY KEY NOT NULL,
# #             ordered_on timestamp NOT NULL,
# #             customer_id int NOT NULL,
# #             delivered_on timestamp,
# #             vendorid int NOT NULL,
# #             status varchar(20) NOT NULL,
# #             bill money NOT NULL
# #     )
# # """)
# # try:
# #     postgres.commit()
# # except Exception as e:
# #     postgres.rollback()
# #     print(e)



# #                       CREATE TABLE ACCOUNTS_RECORD                                    #
# cursor.execute("""
#     create table accounts_record(
#         v_id bigserial PRIMARY KEY NOT NULL references vendors(vendorid),
#         account smallint NOT NULL
#     )
# """)
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)




# #                       CREATE TABLE ORDERsRECORD                                   #
# cursor.execute("""
#     create table ordersrecord(
#         order_id int PRIMARY KEY,
#         vendor_id int NOT NULL references vendors(vendorid),
#         customer_id int NOT NULL references customers(customer_id),
#         ordered_placed_on timestamp NOT NULL,
#         ordered_deliverd_on timestamp ,
#         order_status character(1) CHECK(order_status in('y','n')),
#         description text NOT NULL,
#         amount text NOT NULL
#     )
# """)
# #                       COMMIT AND ROLLBACK IF EXCEPTION                             #

# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)


