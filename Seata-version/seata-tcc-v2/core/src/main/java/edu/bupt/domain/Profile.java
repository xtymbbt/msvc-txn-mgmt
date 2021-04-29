package edu.bupt.domain;

import org.apache.commons.lang.builder.ToStringBuilder;
import org.apache.commons.lang.builder.ToStringStyle;

import java.util.Date;

/**
 * profile对象 profile
 * 
 * @author bridge
 * @date 2021-04-29
 */
public class Profile
{
    private static final long serialVersionUID = 1L;

    /** primary key */
    private Long id;

    /** 现名 */
    private String username;

    /** 年龄 */
    private Integer age;

    /** 性别 */
    private Integer gender;

    /** 群众、共青团员、中共党员、民主人士等 */
    private String politicType;

    /** 出生年月日 */
    private Date birthDate;

    /** 照片 */
    private String photo;

    /** 曾用名 */
    private String oldName;

    /** 身份证号 */
    private String identitiId;

    /** 职业 */
    private String career;

    /** 原籍 */
    private String originHometown;

    /** 出生地 */
    private String birthPlace;

    /** 婚姻状况 */
    private String maritalStatus;

    /** 家庭住址 */
    private String homeAddress;

    /** 现实问题 */
    private String currentProblem;

    /** 教育史 */
    private String educationHistory;

    /** 健康状况 */
    private String healthState;

    /** 详细婚姻状况 */
    private String maritalStatusExplicit;

    /** 就业史 */
    private String employmentHistory;

    /** 以前与社会机构的接触 */
    private String contractInstitusionHistory;

    /** 兴趣爱好 */
    private String hobby;

    /** 其他需要说明的情况 */
    private String otherSituation;

    /** TCC事务状态 */
    private Integer status;

    public void setId(Long id) 
    {
        this.id = id;
    }

    public Long getId() 
    {
        return id;
    }
    public void setUsername(String username) 
    {
        this.username = username;
    }

    public String getUsername() 
    {
        return username;
    }
    public void setAge(Integer age) 
    {
        this.age = age;
    }

    public Integer getAge() 
    {
        return age;
    }
    public void setGender(Integer gender) 
    {
        this.gender = gender;
    }

    public Integer getGender() 
    {
        return gender;
    }
    public void setPoliticType(String politicType) 
    {
        this.politicType = politicType;
    }

    public String getPoliticType() 
    {
        return politicType;
    }
    public void setBirthDate(Date birthDate) 
    {
        this.birthDate = birthDate;
    }

    public Date getBirthDate() 
    {
        return birthDate;
    }
    public void setPhoto(String photo) 
    {
        this.photo = photo;
    }

    public String getPhoto() 
    {
        return photo;
    }
    public void setOldName(String oldName) 
    {
        this.oldName = oldName;
    }

    public String getOldName() 
    {
        return oldName;
    }
    public void setIdentitiId(String identitiId) 
    {
        this.identitiId = identitiId;
    }

    public String getIdentitiId() 
    {
        return identitiId;
    }
    public void setCareer(String career) 
    {
        this.career = career;
    }

    public String getCareer() 
    {
        return career;
    }
    public void setOriginHometown(String originHometown) 
    {
        this.originHometown = originHometown;
    }

    public String getOriginHometown() 
    {
        return originHometown;
    }
    public void setBirthPlace(String birthPlace) 
    {
        this.birthPlace = birthPlace;
    }

    public String getBirthPlace() 
    {
        return birthPlace;
    }
    public void setMaritalStatus(String maritalStatus) 
    {
        this.maritalStatus = maritalStatus;
    }

    public String getMaritalStatus() 
    {
        return maritalStatus;
    }
    public void setHomeAddress(String homeAddress) 
    {
        this.homeAddress = homeAddress;
    }

    public String getHomeAddress() 
    {
        return homeAddress;
    }
    public void setCurrentProblem(String currentProblem) 
    {
        this.currentProblem = currentProblem;
    }

    public String getCurrentProblem() 
    {
        return currentProblem;
    }
    public void setEducationHistory(String educationHistory) 
    {
        this.educationHistory = educationHistory;
    }

    public String getEducationHistory() 
    {
        return educationHistory;
    }
    public void setHealthState(String healthState) 
    {
        this.healthState = healthState;
    }

    public String getHealthState() 
    {
        return healthState;
    }
    public void setMaritalStatusExplicit(String maritalStatusExplicit) 
    {
        this.maritalStatusExplicit = maritalStatusExplicit;
    }

    public String getMaritalStatusExplicit() 
    {
        return maritalStatusExplicit;
    }
    public void setEmploymentHistory(String employmentHistory) 
    {
        this.employmentHistory = employmentHistory;
    }

    public String getEmploymentHistory() 
    {
        return employmentHistory;
    }
    public void setContractInstitusionHistory(String contractInstitusionHistory) 
    {
        this.contractInstitusionHistory = contractInstitusionHistory;
    }

    public String getContractInstitusionHistory() 
    {
        return contractInstitusionHistory;
    }
    public void setHobby(String hobby) 
    {
        this.hobby = hobby;
    }

    public String getHobby() 
    {
        return hobby;
    }
    public void setOtherSituation(String otherSituation) 
    {
        this.otherSituation = otherSituation;
    }

    public String getOtherSituation() 
    {
        return otherSituation;
    }
    public void setStatus(Integer status) 
    {
        this.status = status;
    }

    public Integer getStatus() 
    {
        return status;
    }

    @Override
    public String toString() {
        return new ToStringBuilder(this, ToStringStyle.MULTI_LINE_STYLE)
            .append("id", getId())
            .append("username", getUsername())
            .append("age", getAge())
            .append("gender", getGender())
            .append("politicType", getPoliticType())
            .append("birthDate", getBirthDate())
            .append("photo", getPhoto())
            .append("oldName", getOldName())
            .append("identitiId", getIdentitiId())
            .append("career", getCareer())
            .append("originHometown", getOriginHometown())
            .append("birthPlace", getBirthPlace())
            .append("maritalStatus", getMaritalStatus())
            .append("homeAddress", getHomeAddress())
            .append("currentProblem", getCurrentProblem())
            .append("educationHistory", getEducationHistory())
            .append("healthState", getHealthState())
            .append("maritalStatusExplicit", getMaritalStatusExplicit())
            .append("employmentHistory", getEmploymentHistory())
            .append("contractInstitusionHistory", getContractInstitusionHistory())
            .append("hobby", getHobby())
            .append("otherSituation", getOtherSituation())
            .append("status", getStatus())
            .toString();
    }
}
