package edu.bupt.dao;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.math.BigDecimal;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 8:02
 */
@Mapper
public interface PaymentDao {
    void decrease(@Param("userId") Long userId, @Param("money")BigDecimal money);
}
