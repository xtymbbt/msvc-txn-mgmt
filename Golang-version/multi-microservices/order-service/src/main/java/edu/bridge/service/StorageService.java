package edu.bridge.service;

import edu.bridge.domain.CommonResult;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:11
 */
@FeignClient(value = "storage-service")
public interface StorageService {
    @PostMapping(value = "/storage/decrease")
    CommonResult decrease(@RequestParam("productId") Long productId, @RequestParam("count") Integer count,
                          @RequestParam(value = "UUID", required = false) UUID globalTransactionUUID,
                          @RequestParam(value = "pos", required = false) Integer pos);
}
