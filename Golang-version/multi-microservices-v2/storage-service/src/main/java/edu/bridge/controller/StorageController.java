package edu.bridge.controller;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.service.StorageService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
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

    @PostMapping(value = "/storage/decrease")
    public CommonResult decrease(@RequestParam("productId") Long productId,
                                 @RequestParam("count") Integer count,
                                 @RequestBody(required = false) CommonRequestBody commonRequestBody,
                                 @RequestParam(value = "child", required = false) String child){
        if (commonRequestBody == null) {
            commonRequestBody = new CommonRequestBody(UUID.randomUUID(), "root", "");
        }
        storageService.decrease(productId,count, commonRequestBody, child);
        return new CommonResult(200, "decrease storage succeeded!");
    }
}
