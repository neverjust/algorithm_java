# SpringBoot注解





## @ConfigurationProperties

> 告诉SpringBoot将本类中的所有属性和配置文件中相关的配置进行绑定
>
> 需要导入Configuration Processor 配置文件处理器的包
>
> 只有这个组件是容器中的组件，才能使用容器提供的功能
>
> 用@Component 

```java
//Person.java
@Component
@ConfigurationProperties(prefix = "person")
//prefix = "person" :告诉配置文件中那个下面的所有属性进行一一映射
public class Person {
    private String Name;
    private int age;
    private Boolean Name;
    private String birth;
    //一定要生成每个属性的set方法！！！！
}
```

```yaml
#application.yml
server:
  port: 8081
person:
  Name: 卞光贤
  age: 29
  boss: false
```

```properties
#application.properties
person.Name=卞光贤
peson.age=29
#可能出现乱码 idea是utf-8
```









## @Value

> 等于原来的<bean> <bean/>

```java
//Person.java
@Component
public class Person {
    @Value("${person.name}")
    private String Name;
    @Value(#{10*2})
    private int age;
    @Value("true")
    private Boolean boss;
    private String birth;
    //一定要生成每个属性的set方法！！！！
}
```

| 区别                    |  ConfigurationProperties   |    Value     |
| ----------------------- | :------------------------: | :----------: |
| 功能                    | 批量注入配置文件中的的属性 | 一个一个指定 |
| 松散绑定                |            支持            |    不支持    |
| SpEl                    |           不支持           |     支持     |
| JSR303数据校验          |            支持            |    不支持    |
| 复杂类型(map,list..etc) |            支持            |    不支持    |

> 松散绑定--------驼峰 - _相互支持
>
> SpEl  Spring表达式

**@Value多用于某个业务逻辑中需要获取一下配置文件中的某个配置**

**@ConfigurationProperties适用于多个配置的读取 并进行数据校验**







## @PropertySource

> 加载指定配置文件
>
> @ConfigurationProperties 默认是从配置文件中获取值

```java
@PropertySource(value = {"classpath:person.properties"})
```







## @ImportResource

> 导入 