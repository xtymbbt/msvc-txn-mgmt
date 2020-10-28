package edu.bupt.controller;

import edu.bupt.domain.CommonResult;
import edu.bupt.service.StorageService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 7:03
 */
@RestController
public class StorageController {
    @Autowired
    private StorageService storageService;

    @RequestMapping(value = "/storage/decrease")
    public CommonResult decrease(@RequestParam("productId") Long productId,
                                 @RequestParam("count") Integer count){
        storageService.decrease(productId,count);
        return new CommonResult(200, "decrease storage succeeded!");
    }
}
