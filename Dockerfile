FROM jupyter/scipy-notebook

MAINTAINER Jupyter Project <jupyter@googlegroups.com>

USER $NB_USER

ENV KERAS_BACKEND=tensorflow

# Install Python 3 Tensorflow
#RUN conda install --quiet --yes 'tensorflow=1.0*' keras h5py yaml pydot

# Install Python 2 Tensorflow
RUN conda install --quiet --yes -n python2 tensorflow keras h5py yaml pydot graphviz matplotlib

ADD keras.json /home/jovyan/.keras/
