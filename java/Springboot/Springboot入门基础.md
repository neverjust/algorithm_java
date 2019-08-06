# Springboot 笔记

# 标注

1. @SpringBootApplication

   > 标注在某个雷伤说明这个类是Springboot的主配置类
   >
   > 通过运行这个类的main方法来启动SpringBoot应用
   >
   > ```java
   > //@SpringBootApplication实际上是一个组合注解 
   > @java.lang.annotation.Target({java.lang.annotation.ElementType.TYPE})
   > @java.lang.annotation.Retention(java.lang.annotation.RetentionPolicy.RUNTIME)
   > @java.lang.annotation.Documented
   > @java.lang.annotation.Inherited
   > @org.springframework.boot.SpringBootConfiguration
   > @org.springframework.boot.autoconfigure.EnableAutoConfiguration
   > @org.springframework.context.annotation.ComponentScan(excludeFilters = {@org.springframework.context.annotation.ComponentScan.Filter(type = org.springframework.context.annotation.FilterType.CUSTOM, classes = {org.springframework.boot.context.TypeExcludeFilter.class}), @org.springframework.context.annotation.ComponentScan.Filter(type = org.springframework.context.annotation.FilterType.CUSTOM, classes = {org.springframework.boot.autoconfigure.AutoConfigurationExcludeFilter.class})})
   > 
   > ```
   >
   > 加载各个配置类 组件
   >
   > ![Snip20190314_3](/Users/xian/Desktop/notes/images/Snip20190314_3.png)是

```java
@ResponseBody //将ResponseBody放在这是等于放在所有方法前(如果是对象还会转化为Json数据)
@Controller//这两者等同于 @RestCOntroller
public class helloController {

  @RequestMapping("hello")
  public String hello(){
    return "hi";
  }
}
```

## 开发模式的两个依赖

开发模式的两个依赖 修改之后 自动重启

```xml
<dependency>
            <groupId>org.springframework</groupId>
            <artifactId>springloaded</artifactId>
</dependency>
<dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
</dependency>
```

### 访问静态资源路径

calsspath里面的/static /public或者/META-INF/resources

或者在配置中aplication.properties中配置spring.resources.static-locations=classpath:/static