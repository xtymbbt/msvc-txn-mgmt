package edu.bupt.mapper;

import edu.bupt.entity.Payment;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Param;

import java.math.BigDecimal;

public interface PaymentMapper extends BaseMapper<Payment> {
    void decrease(Long userId, BigDecimal money);

    void updateFrozen(@Param("userId") Long userId, @Param("residue") BigDecimal residue, @Param("frozen") BigDecimal frozen);

    void updateFrozenToUsed(@Param("userId") Long userId, @Param("money") BigDecimal money);

    void updateFrozenToResidue(@Param("userId") Long userId, @Param("money") BigDecimal money);
}