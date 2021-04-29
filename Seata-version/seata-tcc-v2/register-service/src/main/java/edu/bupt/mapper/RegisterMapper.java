package edu.bupt.mapper;

import edu.bupt.domain.Register;
import java.util.List;

/**
 * registerMapper接口
 * 
 * @author bridge
 * @date 2021-04-29
 */
public interface RegisterMapper 
{
    /**
     * 查询register
     * 
     * @param id registerID
     * @return register
     */
    public Register selectRegisterById(Long id);

    /**
     * 查询register列表
     * 
     * @param register register
     * @return register集合
     */
    public List<Register> selectRegisterList(Register register);

    /**
     * 新增register
     * 
     * @param register register
     * @return 结果
     */
    public int insertRegister(Register register);

    /**
     * 修改register
     * 
     * @param register register
     * @return 结果
     */
    public int updateRegister(Register register);

    /**
     * 删除register
     * 
     * @param id registerID
     * @return 结果
     */
    public int deleteRegisterById(Long id);

    /**
     * 批量删除register
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    public int deleteRegisterByIds(String[] ids);
}
