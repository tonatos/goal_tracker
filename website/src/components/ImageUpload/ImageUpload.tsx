import { IoClose } from 'react-icons/io5';
import { ChangeEvent, DragEvent, useEffect, useRef, useState } from 'react';
import styles from './ImageUpload.module.css';
import { useImageUploadActions } from '../../actions/upload.actions';

interface ImageUpload {
    initial?: string,
    backgroundText: string,
    onUpload?: (file: File) => void,
    inputProps?: object,
}

const isFileImage = (file: File) => {
    const acceptedImageTypes = ['image/gif', 'image/jpeg', 'image/png'];
    return file && acceptedImageTypes.includes(file['type'])
}

export const ImageUpload = (props: ImageUpload) => {
    const uploadActions = useImageUploadActions();
    const fileInputRef = useRef<HTMLInputElement | null>(null);
    const [uploadedImage, setUploadedImage] = useState({} as File);
    const [dragActive, setDragActive] = useState(false);

    const setExistsImage = async (url: string) => {
        const image = await uploadActions.download(url);
        const file = new File([image], URL.createObjectURL(image), {type: image.type});
        setUploadedImage(file)
    }

    useEffect(() => {
        if (!props.initial || props.initial == undefined) {
            return;
        }
        setExistsImage(props.initial);
    }, [props.initial]);

    const handleDrag = (e: DragEvent) => {
        e.preventDefault();
        e.stopPropagation();

        if (e.type === 'dragenter' || e.type === 'dragover') {
          setDragActive(true);
        } else if (e.type === 'dragleave') {
          setDragActive(false);
        }
    };

    const handleDrop = (e: DragEvent) => {
      e.preventDefault();
      e.stopPropagation();

      setDragActive(false);

      if (
        !e.dataTransfer.files ||
        !e.dataTransfer.files[0] ||
        !isFileImage(e.dataTransfer.files[0])
      ) {
        return;
      }

      handleFiles(e.dataTransfer.files[0]);
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
      e.preventDefault();
      if (e.target.files && e.target.files[0]) {
        handleFiles(e.target.files[0]);
      }
    };

    const handleRemove = (e: React.MouseEvent<HTMLElement>) => {
        e.preventDefault();
        setUploadedImage({} as File);
    };

    const handleFiles = (file: File) => {
        setUploadedImage(file);
        if (fileInputRef.current) {
            const dt = new DataTransfer();
            dt.items.add(file);

            fileInputRef.current.files = dt.files;
        }
        if (props.onUpload) {
            props.onUpload(file);
        }
    };

    const Dropzone = (
        <>
            <div className={styles.ImageUploadIcon}>
                <svg xmlns="http://www.w3.org/2000/svg" width="55px" height="48px" viewBox="0 0 55 48" version="1.1">
                    <title>Fill 85</title>
                    <defs>
                        <linearGradient x1="90.8601591%" y1="32.9298052%" x2="1.20563282e-14%" y2="75.4375566%" id="linearGradient-1">
                            <stop stopColor="#0DCFAB" stopOpacity="0" offset="0%"/>
                            <stop stopColor="#0DCFAB" offset="100%"/>
                        </linearGradient>
                        <path d="M586,567 L586,567 L590.475551,557.447259 L595.825681,562.543242 L601.921053,553.077578 L610.163867,561.478528 L617.439716,547 L619,567 L586,567 Z M589.489633,545 L589.489633,545 C591.461493,545 593,546.554865 593,548.501035 C593,550.451346 591.461493,552 589.489633,552 C587.613152,552 586,550.451346 586,548.501035 C586,546.554865 587.613152,545 589.489633,545 Z M572.442952,565.428169 L572.442952,531.624567 C572.442952,531.381171 572.678854,531.182398 572.924925,531.182398 L613.073042,531.182398 C613.317078,531.182398 613.504173,531.381171 613.504173,531.624567 L613.504173,533.762392 L617,533.762392 L617,528.527357 C617,528.245424 616.735627,528 616.459052,528 L569.518578,528 C569.244037,528 569,528.245424 569,528.527357 L569,568.484813 C569,568.756604 569.244037,569 569.518578,569 L575.692709,569 L575.692709,565.864252 L572.924925,565.864252 C572.678854,565.864252 572.442952,565.679677 572.442952,565.428169 Z M623.503497,535 L576.510747,535 C576.244182,535 576,535.227158 576,535.525303 L576,575.42602 C576,575.750532 576.244182,576 576.510747,576 L623.503497,576 C623.784306,576 624,575.750532 624,575.42602 L624,535.525303 C624,535.227158 623.784306,535 623.503497,535 L623.503497,535 Z" id="path-2"/>
                    </defs>
                    <g id="Page-1" stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                        <g id="Fill-85" transform="translate(-569.000000, -528.000000)">
                            <use fill="#16B0E3" href="#path-2"/>
                            <use fill="url(#linearGradient-1)" href="#path-2"/>
                        </g>
                    </g>
                </svg>
            </div>
            <div className={styles.ImageUploadText}>
                <h3 className={styles.ImageUploadHeader}>
                    Перетащите сюда картинку вашей цели или <a>выберите ее</a>
                </h3>
                <p className={styles.ImageUploadHelptext}>Подходят png, желательно с прозрачным фоном!</p>
            </div>
        </>
    );

    const Uploaded = (
        <>
            <div className={styles.ImageUploadUploadedContainer}>
                <a onClick={handleRemove} className={styles.ImageUploadRemove}>
                    <IoClose size={24} />
                </a>
                <img src={uploadedImage.lastModified ? URL.createObjectURL(uploadedImage) : undefined} className={styles.ImageUploadImage} />
            </div>
            <h1 className={styles.ImageUploadBackgroundText}>{props.backgroundText}</h1>
        </>
    );
    
    return (
        <label
            className={[
                styles.ImageUpload,
                dragActive ? styles.ImageUploadActive : '',
                uploadedImage.lastModified ? styles.ImageUploadUploaded : ''
            ].join(' ')}
            htmlFor="images"
            onDragEnter={handleDrag}
            onDragLeave={handleDrag}
            onDragOver={handleDrag}
            onDrop={handleDrop}
        >
            {uploadedImage.lastModified ? Uploaded : Dropzone}
            <input
                ref={fileInputRef}
                className={styles.ImageUploadInput}
                type="file"
                id="images"
                accept="image/*"
                onChange={handleChange}
                {...(props.inputProps ? props.inputProps : {})}
            />
        </label>
    );
}