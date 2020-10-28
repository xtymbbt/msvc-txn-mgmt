package edu.bupt.dao;

import edu.bupt.domain.Order;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/26 9:52
 */
@Mapper
public interface OrderDao {
    void create(Order order);
    void update(@Param("userId") Long userId, @Param("status") Integer status);
}
