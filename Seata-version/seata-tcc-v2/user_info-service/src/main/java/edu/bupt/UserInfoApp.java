package edu.bupt;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@EnableFeignClients
@MapperScan("edu.bupt.mapper")
@SpringBootApplication
public class UserInfoApp {

    public static void main(String[] args) {
        SpringApplication.run(UserInfoApp.class, args);
    }

}
