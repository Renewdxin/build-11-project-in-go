
link: [通过构建 11 个项目来学习 Go 编程 – 完整课程](https://www.bilibili.com/video/BV1qV4y1P7a1/?share_source=copy_web&vd_source=e9ba45c7af100743ab57a5fc5f870c41)
youtube link [Learn Go Programming by Building 11 Projects – Full Course](https://youtu.be/jFfo23yIWac?si=A3IT65ln3Fu7FBUf)


# MX、SPF 和 DMARC 记录

MX、SPF 和 DMARC 记录是与域名系统（DNS）相关的记录，用于配置和提供有关邮件服务器、电子邮件发送策略以及邮件身份验证的信息。这些记录对于确保邮件的可靠传递、减少垃圾邮件和提高电子邮件安全性至关重要。

1. MX 记录（Mail Exchange）:
    - 用途：MX 记录指定了接收电子邮件的邮件服务器的地址。
    - 配置：MX 记录由域名所有者在 DNS 中配置，它包含一个或多个邮件服务器的域名和相应的优先级，表示邮件应该按照优先级顺序传递给这些邮件服务器。

2. SPF 记录（Sender Policy Framework）:
    - 用途：SPF 记录用于指定哪些邮件服务器有权发送特定域名的电子邮件。
    - 配置：SPF 记录通过 DNS 配置，其中包含了允许发送邮件的授权邮件服务器的 IP 地址。接收方邮件服务器可以检查 SPF 记录，验证发送方是否是授权的邮件服务器。

3. DMARC 记录（Domain-based Message Authentication, Reporting, and Conformance）:
    - 用途：DMARC 记录用于提供一种机制，确保电子邮件发送方的身份验证，并为域名的所有者提供有关未通过身份验证的邮件的报告。
    - 配置：DMARC 记录由域名所有者在 DNS 中配置。它指定了用于验证电子邮件的 SPF 和 DKIM（DomainKeys Identified Mail）策略，并定义了未通过验证的邮件的处理方式，以及报告的接收地址。

这些记录共同用于帮助确保电子邮件的可靠性、提高邮件的安全性，并减少垃圾邮件和欺诈性电子邮件。在设置和配置这些记录时，域名管理员通常需要使用 DNS 管理界面或控制面板来进行相应的配置。