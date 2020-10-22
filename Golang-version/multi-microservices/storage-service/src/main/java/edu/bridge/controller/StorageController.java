package edu.bridge.controller;

import edu.bridge.domain.CommonResult;
import edu.bridge.service.StorageService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:58
 */
@RestController
public class StorageController {
    @Autowired
    private StorageService storageService;

    @RequestMapping(value = "/storage/decrease")
    public CommonResult decrease(@RequestParam("productId") Long productId,
                                 @RequestParam("count") Integer count,
                                 @RequestParam(value = "UUID", required = false) UUID globalTransactionUUID,
                                 @RequestParam(value = "pos", required = false) Integer pos){
        if (globalTransactionUUID == null) {globalTransactionUUID = UUID.randomUUID();}
        if (pos == null) {pos = 0;}else {pos++;}
        storageService.decrease(productId,count, globalTransactionUUID, pos);
        return new CommonResult(200, "decrease storage succeeded!");
    }
}
