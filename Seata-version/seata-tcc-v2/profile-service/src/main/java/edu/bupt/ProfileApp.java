package edu.bupt;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@MapperScan("edu.bupt.mapper")
@SpringBootApplication
public class ProfileApp {

    public static void main(String[] args) {
        SpringApplication.run(ProfileApp.class, args);
    }

}
