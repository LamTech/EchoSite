30 1 10 * * /usr/bin/certbot renew && kill -9 $(pidof echosited) # 每月10日1点30分执行一次
