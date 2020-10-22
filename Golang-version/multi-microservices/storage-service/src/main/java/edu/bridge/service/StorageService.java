package edu.bridge.service;

import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:43
 */
public interface StorageService {
    void decrease(Long productId, Integer count, UUID uuid, int pos);
}
