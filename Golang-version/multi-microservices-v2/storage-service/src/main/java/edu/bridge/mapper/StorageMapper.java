package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import feign.Param;

import java.util.HashMap;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:42
 */
public interface StorageMapper {
    void decrease(@Param("productId") Long productId, @Param("count") Integer count,
                  CommonRequestBody commonRequestBody, HashMap<String, Boolean> children);
}
