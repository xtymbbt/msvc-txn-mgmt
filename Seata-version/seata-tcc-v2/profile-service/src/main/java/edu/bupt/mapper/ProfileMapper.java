package edu.bupt.mapper;

import edu.bupt.domain.Profile;
import java.util.List;

/**
 * profileMapper接口
 * 
 * @author bridge
 * @date 2021-04-27
 */
public interface ProfileMapper 
{
    /**
     * 查询profile
     * 
     * @param id profileID
     * @return profile
     */
    public Profile selectProfileById(Long id);

    /**
     * 查询profile列表
     * 
     * @param profile profile
     * @return profile集合
     */
    public List<Profile> selectProfileList(Profile profile);

    /**
     * 新增profile
     * 
     * @param profile profile
     * @return 结果
     */
    public int insertProfile(Profile profile);

    /**
     * 修改profile
     * 
     * @param profile profile
     * @return 结果
     */
    public int updateProfile(Profile profile);

    /**
     * 删除profile
     * 
     * @param id profileID
     * @return 结果
     */
    public int deleteProfileById(Long id);

    /**
     * 批量删除profile
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    public int deleteProfileByIds(String[] ids);
}
