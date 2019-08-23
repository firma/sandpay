# sandpay
杉德支付 网银支付&amp;代下单服务

提供的pfx证书在golang中解析会失败 需要转为pem证书解析
  pfx转pem方法,需要安装openssl
  
解析pem证书的公钥和私钥
  openssl pkcs12 -in xxxx.pfx -nodes -out server.pem 生成为原生格式pem 私钥
  openssl rsa -in server.pem -out server.key  生成为rsa格式私钥文件
  openssl x509 -in server.pem  -out server.crt
  
  openssl pkcs12 -in xxxx.pfx -clcerts -nokeys -out key.cert
  
  在目录下生成server.key即为私钥文件.

  在目录下生成server.crt即为公钥文件.
  
