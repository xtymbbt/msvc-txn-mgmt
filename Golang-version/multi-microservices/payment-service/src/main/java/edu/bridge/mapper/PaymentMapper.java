package edu.bridge.mapper;

import java.math.BigDecimal;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:19
 */
public interface PaymentMapper {
    void decrease(Long userId, BigDecimal money, UUID uuid, int pos, boolean isTheLastService);
}
