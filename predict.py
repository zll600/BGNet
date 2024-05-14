import imageio
import torch
import torch.nn.functional as F
import numpy as np
import os, argparse
os.environ["CUDA_VISIBLE_DEVICES"] = "0"
from net.bgnet import Net
from utils.tdataloader import test_dataset

parser = argparse.ArgumentParser()
parser.add_argument('--testsize', type=int, default=416, help='testing size')
parser.add_argument('--pth_path', type=str, default='./checkpoints/best/BGNet.pth')
parser.add_argument('--data_path', type=str, default='./web/data')
parser.add_argument('--save_path', type=str, default='./web/result/')

opt = parser.parse_args()
model = Net()
model.load_state_dict(torch.load(opt.pth_path))
model.cuda()
model.eval()

image_root = '{}/Img/'.format(opt.data_path)
gt_root = '{}/GT/'.format(opt.data_path)
test_loader = test_dataset(image_root, gt_root, opt.testsize)

for i in range(test_loader.size):
        image, gt, name = test_loader.load_data()
        gt = np.asarray(gt, np.float32)
        gt /= (gt.max() + 1e-8)
        image = image.cuda()

        _, _, res, e = model(image)
        res = F.upsample(res, size=gt.shape, mode='bilinear', align_corners=False)
        res = res.sigmoid().data.cpu().numpy().squeeze()
        res = (res - res.min()) / (res.max() - res.min() + 1e-8)
        imageio.imwrite(opt.save_path+name, (res*255).astype(np.uint8))
        # e = F.upsample(e, size=gt.shape, mode='bilinear', align_corners=True)
        # e = e.data.cpu().numpy().squeeze()
        # e = (e - e.min()) / (e.max() - e.min() + 1e-8)
        # imageio.imwrite(save_path+'edge/'+name, (e*255).astype(np.uint8))
