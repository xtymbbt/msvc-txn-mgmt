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
import edu.bupt.domain.Register;
import edu.bupt.service.IRegisterService;

/**
 * register 提供者
 * 
 * @author bridge
 * @date 2021-04-27
 */
@RestController
@RequestMapping("register")
public class RegisterController extends BaseController
{
	
	@Autowired
	private IRegisterService registerService;
	
	/**
	 * 查询${tableComment}
	 */
	@GetMapping("get/{id}")
	public Register get(@PathVariable("id") Long id)
	{
		return registerService.selectRegisterById(id);
		
	}
	
	/**
	 * 查询register列表
	 */
	@GetMapping("list")
	public R list(Register register)
	{
		startPage();
        return result(registerService.selectRegisterList(register));
	}
	
	
	/**
	 * 新增保存register
	 */
	@PostMapping("save")
	public R addSave(@RequestBody Register register)
	{		
		return toAjax(registerService.insertRegister(register));
	}

	/**
	 * 修改保存register
	 */
	@PostMapping("update")
	public R editSave(@RequestBody Register register)
	{		
		return toAjax(registerService.updateRegister(register));
	}
	
	/**
	 * 删除${tableComment}
	 */
	@PostMapping("remove")
	public R remove(String ids)
	{		
		return toAjax(registerService.deleteRegisterByIds(ids));
	}
	
}
