package edu.bupt.mapper;

import edu.bupt.entity.Storage;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Param;

public interface StorageMapper extends BaseMapper<Storage> {
    void decrease(Long productId, Integer count);

    void updateFrozen(@Param("productId") Long productId, @Param("residue") Integer residue, @Param("frozen") Integer frozen);

    void updateFrozenToUsed(@Param("productId") Long productId, @Param("count") Integer count);

    void updateFrozenToResidue(@Param("productId") Long productId, @Param("count") Integer count);
}
