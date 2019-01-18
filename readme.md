以一个简化的业务模型探索 DDD 战术模型的代码实现。

## 业务模型

用于探索实现的业务模型是一个积分账户管理的业务，业务模型如下：

### 开通积分账户

* 开通的积分账户必须属于会员系统中的一个会员。
* 同一个会员允许开通多个积分账户。
* 开通积分账户后账户默认可用积分为零。

### 积分充值

* 可以为指定的账户充值指定额度的积分。
* 充值成功后账户可用积分将增加相应额度。
* 充值积分额度必须是整数不支持小数额度。

### 积分消费

* 可以消费指定账户的指定额度积分。
* 消费成功后账户可用积分将减少相应额度。
* 消费积分必须是整数不支持小数额度。

### 查询账户详情

* 查看账户详情时可查看账户的基本信息和操作记录，包括：
  * 账户所属会员
  * 账户当前可用积分
  * 账户历史充值积分汇总
  * 账户历史消费积分汇总
  * 账户创建时间
  * 账户操作记录
    * 操作时间
    * 操作名称
    * 操作描述


## 探索一: 只使用 DDD

在 base 目录的实现，只使用 OO 和 DDD 的思想探索如何实现 DDD 的战术模型。

## 探索二： DDD + CQRS

basecqrs 目录的实现，在探索一的基础上增加 CQRS 的应用。

## 探索三: DDD + EventSourcing

basees 目录的实现，在探索一的基础上增加 EventSourcing 的应用。

## 控制四: DDD + CQRS + EventSourcing

cqrses 目录的实现，同时使用 CQRS, EventSourcing 的应用。