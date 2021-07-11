# JWT

## JWT的基本原理

JWT 是用来替换 Session 的一种解决方案。因此它不能有大量的计算，必须尽可能的少计算；也不能存储私密的内容。

在设计 JWT 时，需要分成 header、payload、signature 三部分。这三部分都是在后端计算，返回给前端的只是一个 Token 字符串。

header 存储 JWT 元数据。具体而言就是：JWT 是用什麽算法加密的。
```ruby
{
"alg":"sha256",
"typ":"JWT"
}
```

payload 存储具体数据。比如登录用户的 ID。记住，因爲 payload 默认不加密，仅做 base64 编码，爲了安全考虑，尽量不要存太私密的东西。
```ruby
{
"iss":"abc.com", //签发人
"exp":time()+600, //过期时间，10 分钟后
"nbf":time()+2, //生效时间，2 秒后
"iat":time(), //签发时间
"uid":uid //userid 用户 ID
}
```

上面的例子里，最重要的是 exp 和 uid。exp(过期时间) 如果不做限制，一但 JWT 泄漏，任何人都可以用它来登录，永远有效。uid(用户 ID) 是我们用来替代 session，识别用户的信息，也是我们这个 payload 存在的目的。

signature 是签名。它用于保证前两个数据没有被人改过。将前两个数据(header, payload)的 base64 编码 用 "." 连接起来，再进行加密。也仅仅在签名的生成上，用了一次加密算法。
```ruby
HMACSHA256(
base64UrlEncode(header) + "." +
base64UrlEncode(payload),
secret)
```

secret 是我们自定义的密钥。

在上面的三部分生成完毕之后，用 "." 连接起来，传给前端。以后每次请求，都要使用 JWT 来验证身份。因爲 payload 和 header 都不做加密，因此前端传来时，可以反 base64 解开，看信息。最后，再用 签名 验证一下信息是否是僞造的就好了。

## JWT的优点

1. 解决了多服务器之间session共享存在的问题。将用户数据存放在客户端
2. 预防了CSRF攻击，jwt传来的token存在local storage里面设置HttpOnly=false。恶意网站无法获取local storage里面的jwt token

## JWT的缺点

1. **更多的空间占用。**如果将原存在服务端session中的各类信息都放在JWT中保存在客户端，可能造成JWT占用的空间变大，需要考虑cookie的空间限制等因素，如果放在Local Storage，则可能受到XSS攻击。
2. **更不安全。**这里是特指将JWT保存在Local Storage中，然后使用Javascript取出后作为HTTP header发送给服务端的方案。在Local Storage中保存敏感信息并不安全，容易受到跨站脚本攻击，跨站脚本（Cross site script，简称xss）是一种“HTML注入”，由于攻击的脚本多数时候是跨域的，所以称之为“跨域脚本”，这些脚本代码可以盗取cookie或是Local Storage中的数据。可以从这篇文章查看[XSS攻击](https://link.jianshu.com?t=HTTP://www.cnblogs.com/luminji/archive/2012/05/22/2507185.html)的原理解释。
3. **无法作废已颁布的令牌。**所有的认证信息都在JWT中，由于在服务端没有状态，即使你知道了某个JWT被盗取了，你也没有办法将其作废。在JWT过期之前（你绝对应该设置过期时间），你无能为力。

## JWT的解决方案

不再使用Local Storage存储JWT，使用cookie，并且设置HttpOnly=true，这意味着只能由服务端保存以及通过自动回传的cookie取得JWT，以便防御XSS攻击

在JWT的内容中加入一个随机值作为CSRF令牌，由服务端将该CSRF令牌也保存在cookie中，但设置HttpOnly=false，这样前端Javascript代码就可以取得该CSRF令牌，并在请求API时作为HTTP header传回。服务端在认证时，从JWT中取出CSRF令牌与header中获得CSRF令牌比较，从而实现对CSRF攻击的防护

考虑到cookie的空间限制（大约4k左右），在JWT中尽可能只放“够用”的认证信息，其他信息放在数据库，需要时再获取，同时也解决之前提到的数据过期问题