package edu.bridge.domain;

import lombok.Data;

import java.math.BigDecimal;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:12
 */
@Data
public class Order {
    private Long id;
    private Long userId;
    private Long productId;
    private Integer count;
    private BigDecimal money;
    private Integer status; // order state: 1-unpaid, 2-paid
}