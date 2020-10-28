package edu.bupt.service.impl;

import edu.bupt.dao.StorageDao;
import edu.bupt.service.StorageService;
import lombok.extern.slf4j.Slf4j;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 6:56
 */
@Service
@Slf4j
public class StorageServiceImpl implements StorageService {

    private static final Logger LOGGER = LoggerFactory.getLogger(StorageServiceImpl.class);

    @Resource
    private StorageDao storageDao;

    @Override
    public void decrease(Long productId, Integer count) {
        LOGGER.info("------>begin minus storage<-----");
        storageDao.decrease(productId, count);
        LOGGER.info("------>minus storage ended<-----");
    }
}
