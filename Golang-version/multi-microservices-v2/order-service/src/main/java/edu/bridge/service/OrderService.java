package edu.bridge.service;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Order;

import java.util.HashMap;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:12
 */
public interface OrderService {
    void Create(Order order, CommonRequestBody commonRequestBody);
}
