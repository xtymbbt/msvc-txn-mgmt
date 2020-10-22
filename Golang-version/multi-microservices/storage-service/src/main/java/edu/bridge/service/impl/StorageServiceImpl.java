package edu.bridge.service.impl;

import edu.bridge.mapper.StorageMapper;
import edu.bridge.service.StorageService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:44
 */
@Slf4j
@Service
public class StorageServiceImpl implements StorageService {
    @Resource
    private StorageMapper storageMapper;

    @Override
    public void decrease(Long productId, Integer count, UUID uuid, int pos) {
        UUID currentServiceUUID = UUID.randomUUID();
        log.info("------>begin minus storage<-----");
        storageMapper.decrease(productId, count, uuid, currentServiceUUID, 1, 0, pos);
        log.info("------>minus storage ended<-----");
    }
}
