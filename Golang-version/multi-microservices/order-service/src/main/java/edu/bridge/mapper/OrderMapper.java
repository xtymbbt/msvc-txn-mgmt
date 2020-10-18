package edu.bridge.mapper;

import edu.bridge.domain.Order;

import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:39
 */
public interface OrderMapper {
    void create(Order order, UUID uuid, int pos, boolean isTheLastService);
    void update(Long userId, Integer status, UUID uuid, int pos, boolean isTheLastService);
}
