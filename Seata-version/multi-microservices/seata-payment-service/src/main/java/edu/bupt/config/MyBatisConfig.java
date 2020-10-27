package edu.bupt.config;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.context.annotation.Configuration;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/1 21:59
 */
@Configuration
@MapperScan({"edu.bupt.dao"})
public class MyBatisConfig {
}
