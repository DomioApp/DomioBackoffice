# psql setup for connecting from anywhere



> Hi,
>
> I'd like to allow all connections over TCP/IP networks.  So could I just add the following line to the pg_hba.conf in the $PGDATA directory?
> "host   all     *       trust"
>
> The database server is located inside a firewall.
>
>   


    host all all 0.0.0.0/0 trust

or


    host all all 192.168.0.0/24 trust

(or whatever your subnet is in CIDR style network/size notation)

of course, you also need

    listen_address = '*'

in your postgresql.conf so the server is listening on all network interfaces

Personally, I still prefer to use md5 rather than trust and use password 
authentication for LAN connections.
