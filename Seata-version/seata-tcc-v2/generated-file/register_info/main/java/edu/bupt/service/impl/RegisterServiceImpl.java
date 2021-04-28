package edu.bupt.service.impl;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import edu.bupt.mapper.RegisterMapper;
import edu.bupt.domain.Register;
import edu.bupt.service.IRegisterService;
import com.ruoyi.common.core.text.Convert;

/**
 * registerService业务层处理
 * 
 * @author bridge
 * @date 2021-04-27
 */
@Service
public class RegisterServiceImpl implements IRegisterService 
{
    @Autowired
    private RegisterMapper registerMapper;

    /**
     * 查询register
     * 
     * @param id registerID
     * @return register
     */
    @Override
    public Register selectRegisterById(Long id)
    {
        return registerMapper.selectRegisterById(id);
    }

    /**
     * 查询register列表
     * 
     * @param register register
     * @return register
     */
    @Override
    public List<Register> selectRegisterList(Register register)
    {
        return registerMapper.selectRegisterList(register);
    }

    /**
     * 新增register
     * 
     * @param register register
     * @return 结果
     */
    @Override
    public int insertRegister(Register register)
    {
        return registerMapper.insertRegister(register);
    }

    /**
     * 修改register
     * 
     * @param register register
     * @return 结果
     */
    @Override
    public int updateRegister(Register register)
    {
        return registerMapper.updateRegister(register);
    }

    /**
     * 删除register对象
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    @Override
    public int deleteRegisterByIds(String ids)
    {
        return registerMapper.deleteRegisterByIds(Convert.toStrArray(ids));
    }

    /**
     * 删除register信息
     * 
     * @param id registerID
     * @return 结果
     */
    public int deleteRegisterById(Long id)
    {
        return registerMapper.deleteRegisterById(id);
    }
}
