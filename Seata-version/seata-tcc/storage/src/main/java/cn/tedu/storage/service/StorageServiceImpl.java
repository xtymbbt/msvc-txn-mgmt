package cn.tedu.storage.service;

import cn.tedu.storage.tcc.StorageTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class StorageServiceImpl implements StorageService {
    // @Autowired
    // private StorageMapper storageMapper;

    @Autowired
    private StorageTccAction storageTccAction;

    @Override
    public void decrease(Long productId, Integer count) throws Exception {
        // storageMapper.decrease(productId,count);
        storageTccAction.prepareDecreaseStorage(null, productId, count);
    }

}
