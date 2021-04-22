package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Order;

import java.util.HashMap;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:39
 */
public interface OrderMapper {
    void create(Order order, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children);
    void update(Long userId, Integer status, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children);
}
