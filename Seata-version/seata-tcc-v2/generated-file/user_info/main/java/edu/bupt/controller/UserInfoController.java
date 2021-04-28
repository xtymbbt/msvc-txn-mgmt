package edu.bupt.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.ruoyi.common.core.domain.R;
import com.ruoyi.common.core.controller.BaseController;
import edu.bupt.domain.UserInfo;
import edu.bupt.service.IUserInfoService;

/**
 * user_info 提供者
 * 
 * @author bridge
 * @date 2021-04-27
 */
@RestController
@RequestMapping("user_info")
public class UserInfoController extends BaseController
{
	
	@Autowired
	private IUserInfoService userInfoService;
	
	/**
	 * 查询${tableComment}
	 */
	@GetMapping("get/{id}")
	public UserInfo get(@PathVariable("id") Long id)
	{
		return userInfoService.selectUserInfoById(id);
		
	}
	
	/**
	 * 查询user_info列表
	 */
	@GetMapping("list")
	public R list(UserInfo userInfo)
	{
		startPage();
        return result(userInfoService.selectUserInfoList(userInfo));
	}
	
	
	/**
	 * 新增保存user_info
	 */
	@PostMapping("save")
	public R addSave(@RequestBody UserInfo userInfo)
	{		
		return toAjax(userInfoService.insertUserInfo(userInfo));
	}

	/**
	 * 修改保存user_info
	 */
	@PostMapping("update")
	public R editSave(@RequestBody UserInfo userInfo)
	{		
		return toAjax(userInfoService.updateUserInfo(userInfo));
	}
	
	/**
	 * 删除${tableComment}
	 */
	@PostMapping("remove")
	public R remove(String ids)
	{		
		return toAjax(userInfoService.deleteUserInfoByIds(ids));
	}
	
}
